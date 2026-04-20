import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_test/flutter_test.dart';

import 'package:groupevents/app.dart';

void main() {
  testWidgets('home screen renders starter descriptor', (
    WidgetTester tester,
  ) async {
    await tester.pumpWidget(const ProviderScope(child: GroupEventsApp()));

    expect(find.text('Group Events'), findsOneWidget);
    expect(find.text('Server-driven home starter'), findsOneWidget);
    expect(find.text('Descriptor seam is the feature seam'), findsOneWidget);
    expect(find.text('Refresh descriptor'), findsOneWidget);
  });
}
