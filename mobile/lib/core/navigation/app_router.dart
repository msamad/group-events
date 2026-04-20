import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:sdui/sdui.dart';

import '../../features/home/presentation/screens/home_screen.dart';

final appRouterProvider = Provider<AppRouter>((ref) => AppRouter());

class AppRouter {
  static const homeRoute = '/';
  static const sduiRoute = '/screen';

  Route<dynamic> onGenerateRoute(RouteSettings settings) {
    switch (settings.name) {
      case sduiRoute:
        final descriptor = settings.arguments;

        return MaterialPageRoute<void>(
          builder: (_) => HomeScreen(
            descriptor: descriptor is SduiDescriptor ? descriptor : null,
          ),
          settings: settings,
        );
      case homeRoute:
      default:
        return MaterialPageRoute<void>(
          builder: (_) => const HomeScreen(),
          settings: settings,
        );
    }
  }
}
