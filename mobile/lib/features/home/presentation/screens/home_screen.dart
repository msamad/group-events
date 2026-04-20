import 'package:flutter/material.dart';
import 'package:sdui/sdui.dart';

import '../../../../shared/widgets/status_panel.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key, this.descriptor});

  final SduiDescriptor? descriptor;

  static final SduiDescriptor
  _starterDescriptor = SduiDescriptor.fromJson(const <String, Object?>{
    'screenId': 'home',
    'title': 'Server-driven home starter',
    'subtitle':
        'This shell renders descriptor payloads from the shared SDUI package so group-specific behavior can stay on the server contract.',
    'metadata': <String, Object?>{'source': 'starter-fixture'},
    'components': <Map<String, Object?>>[
      <String, Object?>{
        'id': 'seam',
        'type': 'info_panel',
        'title': 'Descriptor seam is the feature seam',
        'body':
            'The mobile client owns shell layout and action plumbing, while role-aware visibility and workflows arrive in payloads.',
        'props': <String, Object?>{'emphasis': 'positive'},
      },
      <String, Object?>{
        'id': 'baseline',
        'type': 'bullet_list',
        'title': 'Starter baseline',
        'props': <String, Object?>{
          'items': <String>[
            'Shared models parse JSON descriptors into immutable Dart objects.',
            'The renderer supports nested content blocks and backend-defined actions.',
            'Feature screens can swap fixtures for API payloads without local role branching.',
          ],
        },
      },
      <String, Object?>{
        'id': 'next',
        'type': 'stack',
        'title': 'Next contract slice',
        'children': <Map<String, Object?>>[
          <String, Object?>{
            'id': 'events',
            'type': 'section',
            'title': 'Event list descriptors',
            'body':
                'Backend responses can next supply cards, empty states, and allowed actions for each group membership context.',
          },
          <String, Object?>{
            'id': 'announcements',
            'type': 'section',
            'title': 'Acknowledgements and reactions',
            'body':
                'Read-only announcement groups can become server-defined prompts instead of bespoke UI branches.',
            'actions': <Map<String, Object?>>[
              <String, Object?>{
                'id': 'preview-reaction-flow',
                'label': 'Preview reaction affordance',
                'type': 'noop',
              },
            ],
          },
        ],
      },
    ],
    'actions': <Map<String, Object?>>[
      <String, Object?>{
        'id': 'refresh-home',
        'label': 'Refresh descriptor',
        'type': 'refresh',
      },
      <String, Object?>{
        'id': 'await-backend',
        'label': 'Await live API payload',
        'type': 'noop',
        'payload': <String, Object?>{'reason': 'Backend contract pending'},
      },
    ],
  });

  @override
  Widget build(BuildContext context) {
    final resolvedDescriptor = descriptor ?? _starterDescriptor;

    return Scaffold(
      appBar: AppBar(title: const Text('Group Events')),
      body: SafeArea(
        child: ListView(
          padding: const EdgeInsets.all(20),
          children: [
            const StatusPanel(
              eyebrow: 'Mobile scaffold',
              title:
                  'Flutter shell is ready for descriptor-driven feature work.',
              description:
                  'Riverpod owns app wiring, and the home route now renders a parsed SDUI descriptor so future event and acknowledgement flows can stay contract-first.',
            ),
            const SizedBox(height: 16),
            SduiRenderer(
              descriptor: resolvedDescriptor,
              onAction: (action) {
                final routeLabel = action.route == null
                    ? ''
                    : ' -> ${action.route}';

                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(
                    content: Text(
                      'Action queued: ${action.label} (${action.type})$routeLabel',
                    ),
                  ),
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}
