import 'dart:ffi';
import 'dart:io';
import 'package:ffi/ffi.dart';

typedef StartWireGuardC = Pointer<Utf8> Function(Pointer<Utf8> configPath);
typedef StartWireGuardDart = Pointer<Utf8> Function(Pointer<Utf8> configPath);

typedef StopWireGuardC = Pointer<Utf8> Function(Pointer<Utf8> configPath);
typedef StopWireGuardDart = Pointer<Utf8> Function(Pointer<Utf8> configPath);

typedef SyncWireGuardC = Pointer<Utf8> Function(Pointer<Utf8> configPath);
typedef SyncWireGuardDart = Pointer<Utf8> Function(Pointer<Utf8> configPath);

final DynamicLibrary wireguardLib = Platform.isLinux
    ? DynamicLibrary.open('libwg_quick.so')
    : DynamicLibrary.process();

final StartWireGuardDart startWireGuard = wireguardLib
    .lookup<NativeFunction<StartWireGuardC>>('StartWireGuard')
    .asFunction();

final StopWireGuardDart stopWireGuard = wireguardLib
    .lookup<NativeFunction<StopWireGuardC>>('StopWireGuard')
    .asFunction();

final SyncWireGuardDart syncWireGuard = wireguardLib
    .lookup<NativeFunction<SyncWireGuardC>>('SyncWireGuard')
    .asFunction();
