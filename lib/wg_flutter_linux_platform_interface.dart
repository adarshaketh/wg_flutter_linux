import 'package:plugin_platform_interface/plugin_platform_interface.dart';
import 'wg_flutter_linux_method_channel.dart';

abstract class WgFlutterLinuxPlatform extends PlatformInterface {
  WgFlutterLinuxPlatform() : super(token: _token);

  static final Object _token = Object();

  static WgFlutterLinuxPlatform _instance = MethodChannelWgFlutterLinux();

  static WgFlutterLinuxPlatform get instance => _instance;

  static set instance(WgFlutterLinuxPlatform instance) {
    PlatformInterface.verifyToken(instance, _token);
    _instance = instance;
  }

  Future<String> startWireGuard(String configPath) {
    throw UnimplementedError('startWireGuard() has not been implemented.');
  }

  Future<String> stopWireGuard(String configPath) {
    throw UnimplementedError('stopWireGuard() has not been implemented.');
  }

  Future<String> syncWireGuard(String configPath) {
    throw UnimplementedError('syncWireGuard() has not been implemented.');
  }
}
