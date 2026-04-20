import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:sdui/sdui.dart';

void main() {
  test('parses descriptor payloads into serializable models', () {
    final descriptor = SduiDescriptor.fromJson(const <String, Object?>{
      'screenId': 'home',
      'title': 'Home',
      'subtitle': 'Server payload',
      'components': <Map<String, Object?>>[
        <String, Object?>{
          'id': 'intro',
          'type': 'section',
          'title': 'Intro',
          'body': 'Rendered from JSON',
        },
      ],
      'actions': <Map<String, Object?>>[
        <String, Object?>{
          'id': 'refresh',
          'label': 'Refresh',
          'type': 'refresh',
          'route': '/screen',
        },
      ],
    });

    expect(descriptor.screenId, 'home');
    expect(descriptor.components.single.title, 'Intro');
    expect(descriptor.actions.single.route, '/screen');
    expect(descriptor.toJson()['title'], 'Home');
  });

  test('engine reports supported starter components', () {
    const engine = SduiEngine();

    expect(
      engine.supportsComponent(const SduiComponent(id: 'a', type: 'stack')),
      isTrue,
    );
    expect(
      engine.supportsComponent(const SduiComponent(id: 'b', type: 'unknown')),
      isFalse,
    );
  });

  testWidgets('renderer shows nested components and forwards actions', (
    WidgetTester tester,
  ) async {
    SduiAction? tappedAction;

    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          body: SduiRenderer(
            descriptor: const SduiDescriptor(
              screenId: 'home',
              title: 'Home',
              subtitle: 'Subtitle',
              components: <SduiComponent>[
                SduiComponent(
                  id: 'list',
                  type: 'bullet_list',
                  title: 'Checklist',
                  props: <String, Object?>{
                    'items': <String>['Alpha', 'Beta'],
                  },
                ),
                SduiComponent(
                  id: 'stack',
                  type: 'stack',
                  title: 'Nested',
                  children: <SduiComponent>[
                    SduiComponent(
                      id: 'child',
                      type: 'section',
                      title: 'Child section',
                      actions: <SduiAction>[
                        SduiAction(
                          id: 'preview',
                          label: 'Preview',
                          type: 'noop',
                        ),
                      ],
                    ),
                  ],
                ),
              ],
              actions: <SduiAction>[
                SduiAction(id: 'refresh', label: 'Refresh', type: 'refresh'),
              ],
            ),
            onAction: (action) {
              tappedAction = action;
            },
          ),
        ),
      ),
    );

    expect(find.text('Checklist'), findsOneWidget);
    expect(find.text('Alpha'), findsOneWidget);
    expect(find.text('Child section'), findsOneWidget);

    await tester.tap(find.text('Preview'));
    await tester.pump();

    expect(tappedAction?.id, 'preview');
  });
}
