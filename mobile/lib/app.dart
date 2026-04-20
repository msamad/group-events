import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import 'core/navigation/app_router.dart';
import 'core/theme/app_theme.dart';

class GroupEventsApp extends ConsumerWidget {
  const GroupEventsApp({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final router = ref.watch(appRouterProvider);

    return MaterialApp(
      title: 'Group Events',
      theme: AppTheme.theme,
      onGenerateRoute: router.onGenerateRoute,
      initialRoute: AppRouter.homeRoute,
    );
  }
}
