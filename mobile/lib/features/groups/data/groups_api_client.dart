import 'dart:convert';

import 'package:http/http.dart' as http;

import '../../../core/config/app_config.dart';
import '../models/group.dart';

class GroupsApiClient {
  GroupsApiClient({http.Client? httpClient, String? baseUrl})
    : _httpClient = httpClient ?? http.Client(),
      _baseUrl = baseUrl ?? AppConfig.apiBaseUrl;

  final http.Client _httpClient;
  final String _baseUrl;

  Future<List<Group>> listGroups({int limit = 20, int offset = 0}) async {
    final uri = Uri.parse(
      '$_baseUrl/api/v1/groups?limit=$limit&offset=$offset',
    );
    final response = await _httpClient.get(uri);
    _expectStatus(response, 200);

    final body = _decodeObject(response.body);
    final data = body['data'];
    final ui = body['ui'];

    if (data is! List) {
      return const <Group>[];
    }

    return data
        .whereType<Map>()
        .map(
          (entry) => Group.fromJson(<String, Object?>{
            ..._normalizeMap(entry),
            if (ui is Map) 'ui': _normalizeMap(ui),
          }),
        )
        .toList(growable: false);
  }

  Future<Group> createGroup({
    required String name,
    required String slug,
    String description = '',
  }) async {
    final uri = Uri.parse('$_baseUrl/api/v1/groups');
    final response = await _httpClient.post(
      uri,
      headers: const {'Content-Type': 'application/json'},
      body: jsonEncode(<String, Object?>{
        'name': name,
        'slug': slug,
        'description': description,
      }),
    );

    _expectStatus(response, 201);
    return Group.fromJson(_decodeObject(response.body));
  }

  Future<Group> getGroup(String id) async {
    final uri = Uri.parse('$_baseUrl/api/v1/groups/$id');
    final response = await _httpClient.get(uri);
    _expectStatus(response, 200);

    return Group.fromJson(_decodeObject(response.body));
  }

  Future<Group> updateGroup({
    required String id,
    required String name,
    required String slug,
    String description = '',
  }) async {
    final uri = Uri.parse('$_baseUrl/api/v1/groups/$id');
    final response = await _httpClient.put(
      uri,
      headers: const {'Content-Type': 'application/json'},
      body: jsonEncode(<String, Object?>{
        'name': name,
        'slug': slug,
        'description': description,
      }),
    );

    _expectStatus(response, 200);
    return Group.fromJson(_decodeObject(response.body));
  }

  Future<void> deleteGroup(String id) async {
    final uri = Uri.parse('$_baseUrl/api/v1/groups/$id');
    final response = await _httpClient.delete(uri);
    _expectStatus(response, 204);
  }

  Map<String, Object?> _decodeObject(String rawBody) {
    final decoded = jsonDecode(rawBody);
    if (decoded is Map<String, Object?>) {
      return decoded;
    }
    if (decoded is Map) {
      return _normalizeMap(decoded);
    }
    throw const FormatException('Expected object JSON payload');
  }

  void _expectStatus(http.Response response, int expectedStatus) {
    if (response.statusCode != expectedStatus) {
      throw HttpException(
        statusCode: response.statusCode,
        body: response.body,
        expectedStatus: expectedStatus,
      );
    }
  }
}

class HttpException implements Exception {
  const HttpException({
    required this.statusCode,
    required this.body,
    required this.expectedStatus,
  });

  final int statusCode;
  final String body;
  final int expectedStatus;

  @override
  String toString() {
    return 'HttpException(statusCode: $statusCode, expected: $expectedStatus, body: $body)';
  }
}

Map<String, Object?> _normalizeMap(Map<dynamic, dynamic> value) {
  return value.map((key, entry) => MapEntry(key.toString(), entry));
}
