import 'dart:ffi';
import 'package:ffi/ffi.dart';
import 'wireguard.dart';

class WgFlutterLinux {
  static Future<String> startWireguard(String configPath) async {
    final Pointer<Utf8> configPathUtf8 = configPath.toNativeUtf8();
    final Pointer<Utf8> result = startWireGuard(configPathUtf8);
    malloc.free(configPathUtf8);
    if (result == nullptr) {
      return 'WireGuard started successfully';
    } else {
      final String error = result.cast<Utf8>().toDartString();
      malloc.free(result);
      throw Exception('Failed to start WireGuard: $error');
    }
  }

  static Future<String> stopWireguard(String configPath) async {
    final Pointer<Utf8> configPathUtf8 = configPath.toNativeUtf8();
    final Pointer<Utf8> result = stopWireGuard(configPathUtf8);
    malloc.free(configPathUtf8);
    if (result == nullptr) {
      return 'WireGuard stopped successfully';
    } else {
      final String error = result.cast<Utf8>().toDartString();
      malloc.free(result);
      throw Exception('Failed to stop WireGuard: $error');
    }
  }

  static Future<String> syncWireguard(String configPath) async {
    final Pointer<Utf8> configPathUtf8 = configPath.toNativeUtf8();
    final Pointer<Utf8> result = syncWireGuard(configPathUtf8);
    malloc.free(configPathUtf8);
    if (result == nullptr) {
      return 'WireGuard synced successfully';
    } else {
      final String error = result.cast<Utf8>().toDartString();
      malloc.free(result);
      throw Exception('Failed to sync WireGuard: $error');
    }
  }
}
