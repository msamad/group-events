import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:sdui/sdui.dart';

void main() {
  test('parses descriptor payloads into serializable models', () {
    final descriptor = SduiDescriptor.fromJson(const <String, Object?>{
      'screen': 'home',
      'title': 'Home',
      'subtitle': 'Server payload',
      'components': <Map<String, Object?>>[
        <String, Object?>{
          'id': 'intro',
          'type': 'section',
          'visible': true,
          'title': 'Intro',
          'body': 'Rendered from JSON',
        },
      ],
      'actions': <Map<String, Object?>>[
        <String, Object?>{
          'id': 'refresh',
          'label': 'Refresh',
          'type': 'refresh',
          'endpoint': '/screen',
          'visible': true,
        },
      ],
    });

    expect(descriptor.screen, 'home');
    expect(descriptor.components.single.title, 'Intro');
    expect(descriptor.actions.single.endpoint, '/screen');
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
              screen: 'home',
              title: 'Home',
              subtitle: 'Subtitle',
              components: <SduiComponent>[
                SduiComponent(
                  id: 'list',
                  type: 'bullet_list',
                  visible: true,
                  title: 'Checklist',
                  data: <String, Object?>{
                    'items': <String>['Alpha', 'Beta'],
                  },
                ),
                SduiComponent(
                  id: 'stack',
                  type: 'stack',
                  visible: true,
                  title: 'Nested',
                  children: <SduiComponent>[
                    SduiComponent(
                      id: 'child',
                      type: 'section',
                      visible: true,
                      title: 'Child section',
                      actions: <SduiAction>[
                        SduiAction(
                          id: 'preview',
                          label: 'Preview',
                          type: 'noop',
                          visible: true,
                        ),
                      ],
                    ),
                  ],
                ),
              ],
              actions: <SduiAction>[
                SduiAction(
                  id: 'refresh',
                  label: 'Refresh',
                  type: 'refresh',
                  visible: true,
                ),
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
