import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:sdui/sdui.dart';

void main() {
  testWidgets('fixture descriptor renders expected component types', (
    WidgetTester tester,
  ) async {
    final fixture = File(
      'test/fixtures/sdui/groups_list.json',
    ).readAsStringSync();
    final payload = jsonDecode(fixture) as Map<String, Object?>;

    final descriptor = SduiDescriptor.fromJson(payload);

    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(body: SduiRenderer(descriptor: descriptor)),
      ),
    );

    expect(find.text('Groups'), findsOneWidget);
    expect(find.text('Create event'), findsOneWidget);
    expect(find.textContaining('Unsupported component'), findsNothing);
  });
}
