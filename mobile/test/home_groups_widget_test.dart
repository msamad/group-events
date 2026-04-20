import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:groupevents/features/groups/data/groups_api_client.dart';
import 'package:groupevents/features/groups/models/group.dart';
import 'package:mocktail/mocktail.dart';

import 'mocks/mock_groups_api_client.dart';

class _GroupsPreview extends StatelessWidget {
  const _GroupsPreview({required this.client});

  final GroupsApiClient client;

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<List<Group>>(
      future: client.listGroups(),
      builder: (context, snapshot) {
        if (!snapshot.hasData) {
          return const CircularProgressIndicator();
        }

        return Text('Loaded ${snapshot.data!.length} groups');
      },
    );
  }
}

void main() {
  testWidgets('mock GroupsApiClient drives widget rendering', (tester) async {
    final mockClient = MockGroupsApiClient();
    when(
      () => mockClient.listGroups(
        limit: any(named: 'limit'),
        offset: any(named: 'offset'),
      ),
    ).thenAnswer(
      (_) async => [
        Group.fromJson(const <String, Object?>{
          'id': 'group-1',
          'slug': 'platform',
          'name': 'Platform',
          'createdAt': '2026-04-20T00:00:00Z',
        }),
      ],
    );

    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(body: _GroupsPreview(client: mockClient)),
      ),
    );

    await tester.pumpAndSettle();

    expect(find.text('Loaded 1 groups'), findsOneWidget);
    verify(
      () => mockClient.listGroups(
        limit: any(named: 'limit'),
        offset: any(named: 'offset'),
      ),
    ).called(1);
  });
}
