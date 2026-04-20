import 'package:flutter_test/flutter_test.dart';
import 'package:groupevents/features/groups/models/group.dart';

void main() {
  test('Group.fromJson parses fields including ui', () {
    final group = Group.fromJson(const <String, Object?>{
      'data': <String, Object?>{
        'id': 'group-1',
        'slug': 'platform',
        'name': 'Platform',
        'description': 'Core team',
        'archived': false,
        'createdAt': '2026-04-20T00:00:00Z',
      },
      'ui': <String, Object?>{
        'screenId': 'groups:list',
        'title': 'Groups',
        'components': <Map<String, Object?>>[],
        'actions': <Map<String, Object?>>[],
      },
    });

    expect(group.id, 'group-1');
    expect(group.slug, 'platform');
    expect(group.name, 'Platform');
    expect(group.ui, isNotNull);
  });

  test('Group.fromJson handles missing ui', () {
    final group = Group.fromJson(const <String, Object?>{
      'id': 'group-2',
      'slug': 'ops',
      'name': 'Ops',
      'createdAt': '2026-04-20T00:00:00Z',
    });

    expect(group.ui, isNull);
  });

  test('Group.toJson serializes all model fields', () {
    final source = Group.fromJson(const <String, Object?>{
      'id': 'group-3',
      'slug': 'eng',
      'name': 'Engineering',
      'description': 'Desc',
      'archived': false,
      'createdAt': '2026-04-20T00:00:00Z',
      'ui': <String, Object?>{
        'screenId': 'groups:detail',
        'title': 'Group details',
      },
    });

    final json = source.toJson();

    expect(json['id'], 'group-3');
    expect(json['slug'], 'eng');
    expect(json['name'], 'Engineering');
    expect(json['ui'], isA<Map<String, Object?>>());
  });
}
