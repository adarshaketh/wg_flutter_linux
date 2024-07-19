#include "include/wg_flutter_linux/wg_flutter_linux_plugin.h"

#include <flutter_linux/flutter_linux.h>
#include <gtk/gtk.h>
#include <sys/utsname.h>

#include <cstring>

#include "wg_flutter_linux_plugin_private.h"

#define WG_FLUTTER_LINUX_PLUGIN(obj) \
  (G_TYPE_CHECK_INSTANCE_CAST((obj), wg_flutter_linux_plugin_get_type(), \
                              WgFlutterLinuxPlugin))

struct _WgFlutterLinuxPlugin {
  GObject parent_instance;
};

G_DEFINE_TYPE(WgFlutterLinuxPlugin, wg_flutter_linux_plugin, g_object_get_type())

extern "C" {
    char* StartWireGuard(const char* configPath);
    char* StopWireGuard(const char* configPath);
    char* SyncWireGuard(const char* configPath);
}

// Called when a method call is received from Flutter.
static void wg_flutter_linux_plugin_handle_method_call(
    WgFlutterLinuxPlugin* self,
    FlMethodCall* method_call) {
  g_autoptr(FlMethodResponse) response = nullptr;

  const gchar* method = fl_method_call_get_name(method_call);
  FlValue* args = fl_method_call_get_args(method_call);
  const gchar* configPath = fl_value_get_string(fl_value_lookup_string(args, "configPath"));

  if (strcmp(method, "startWireGuard") == 0) {
    char* result = StartWireGuard(configPath);
    if (result == nullptr) {
      response = FL_METHOD_RESPONSE(fl_method_success_response_new(fl_value_new_string("WireGuard started successfully")));
    } else {
      response = FL_METHOD_RESPONSE(fl_method_error_response_new("START_ERROR", result, nullptr));
      g_free(result);
    }
  } else if (strcmp(method, "stopWireGuard") == 0) {
    char* result = StopWireGuard(configPath);
    if (result == nullptr) {
      response = FL_METHOD_RESPONSE(fl_method_success_response_new(fl_value_new_string("WireGuard stopped successfully")));
    } else {
      response = FL_METHOD_RESPONSE(fl_method_error_response_new("STOP_ERROR", result, nullptr));
      g_free(result);
    }
  } else if (strcmp(method, "syncWireGuard") == 0) {
    char* result = SyncWireGuard(configPath);
    if (result == nullptr) {
      response = FL_METHOD_RESPONSE(fl_method_success_response_new(fl_value_new_string("WireGuard synced successfully")));
    } else {
      response = FL_METHOD_RESPONSE(fl_method_error_response_new("SYNC_ERROR", result, nullptr));
      g_free(result);
    }
  } else if (strcmp(method, "getPlatformVersion") == 0) {
    response = get_platform_version();
  } else {
    response = FL_METHOD_RESPONSE(fl_method_not_implemented_response_new());
  }

  fl_method_call_respond(method_call, response, nullptr);
}

FlMethodResponse* get_platform_version() {
  struct utsname uname_data = {};
  uname(&uname_data);
  g_autofree gchar *version = g_strdup_printf("Linux %s", uname_data.version);
  g_autoptr(FlValue) result = fl_value_new_string(version);
  return FL_METHOD_RESPONSE(fl_method_success_response_new(result));
}

static void wg_flutter_linux_plugin_dispose(GObject* object) {
  G_OBJECT_CLASS(wg_flutter_linux_plugin_parent_class)->dispose(object);
}

static void wg_flutter_linux_plugin_class_init(WgFlutterLinuxPluginClass* klass) {
  G_OBJECT_CLASS(klass)->dispose = wg_flutter_linux_plugin_dispose;
}

static void wg_flutter_linux_plugin_init(WgFlutterLinuxPlugin* self) {}

static void method_call_cb(FlMethodChannel* channel, FlMethodCall* method_call,
                           gpointer user_data) {
  WgFlutterLinuxPlugin* plugin = WG_FLUTTER_LINUX_PLUGIN(user_data);
  wg_flutter_linux_plugin_handle_method_call(plugin, method_call);
}

void wg_flutter_linux_plugin_register_with_registrar(FlPluginRegistrar* registrar) {
  WgFlutterLinuxPlugin* plugin = WG_FLUTTER_LINUX_PLUGIN(
      g_object_new(wg_flutter_linux_plugin_get_type(), nullptr));

  g_autoptr(FlStandardMethodCodec) codec = fl_standard_method_codec_new();
  g_autoptr(FlMethodChannel) channel =
      fl_method_channel_new(fl_plugin_registrar_get_messenger(registrar),
                            "wg_flutter_linux",
                            FL_METHOD_CODEC(codec));
  fl_method_channel_set_method_call_handler(channel, method_call_cb,
                                            g_object_ref(plugin),
                                            g_object_unref);

  g_object_unref(plugin);
}
