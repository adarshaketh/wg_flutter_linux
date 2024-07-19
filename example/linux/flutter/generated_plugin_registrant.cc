//
//  Generated file. Do not edit.
//

// clang-format off

#include "generated_plugin_registrant.h"

#include <wg_flutter_linux/wg_flutter_linux_plugin.h>

void fl_register_plugins(FlPluginRegistry* registry) {
  g_autoptr(FlPluginRegistrar) wg_flutter_linux_registrar =
      fl_plugin_registry_get_registrar_for_plugin(registry, "WgFlutterLinuxPlugin");
  wg_flutter_linux_plugin_register_with_registrar(wg_flutter_linux_registrar);
}
