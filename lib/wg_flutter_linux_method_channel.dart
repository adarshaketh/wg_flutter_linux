import 'package:flutter/foundation.dart';
import 'package:flutter/services.dart';
import 'wg_flutter_linux_platform_interface.dart';
import 'wg_flutter_linux.dart';

class MethodChannelWgFlutterLinux extends WgFlutterLinuxPlatform {
  @visibleForTesting
  final methodChannel = const MethodChannel('wg_flutter_linux');

  @override
  Future<String> startWireGuard(String configPath) async {
    final String result = await WgFlutterLinux.startWireguard(configPath);
    return result;
  }

  @override
  Future<String> stopWireGuard(String configPath) async {
    final String result = await WgFlutterLinux.stopWireguard(configPath);
    return result;
  }

  @override
  Future<String> syncWireGuard(String configPath) async {
    final String result = await WgFlutterLinux.syncWireguard(configPath);
    return result;
  }
}
