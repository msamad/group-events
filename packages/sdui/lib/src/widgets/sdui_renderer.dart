import 'package:flutter/material.dart';

import '../engine/sdui_engine.dart';
import '../models/sdui_action.dart';
import '../models/sdui_component.dart';
import '../models/sdui_descriptor.dart';

class SduiRenderer extends StatelessWidget {
  const SduiRenderer({
    super.key,
    required this.descriptor,
    this.engine = const SduiEngine(),
    this.onAction,
  });

  final SduiDescriptor descriptor;
  final SduiEngine engine;
  final ValueChanged<SduiAction>? onAction;

  @override
  Widget build(BuildContext context) {
    final components = engine
        .resolveComponents(descriptor)
        .where((component) => component.visible)
        .toList(growable: false);
    final actions = engine
        .resolveActions(descriptor)
        .where((action) => action.visible)
        .toList(growable: false);
    final theme = Theme.of(context);

    return Card(
      child: Padding(
        padding: const EdgeInsets.all(20),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              descriptor.title,
              style: theme.textTheme.titleLarge?.copyWith(
                fontWeight: FontWeight.w700,
              ),
            ),
            if (descriptor.subtitle != null) ...[
              const SizedBox(height: 8),
              Text(
                descriptor.subtitle!,
                style: theme.textTheme.bodyLarge?.copyWith(
                  color: theme.colorScheme.onSurfaceVariant,
                  height: 1.4,
                ),
              ),
            ],
            const SizedBox(height: 16),
            for (final component in components) ...[
              _SduiComponentNode(
                component: component,
                engine: engine,
                onAction: onAction,
              ),
              const SizedBox(height: 12),
            ],
            if (actions.isNotEmpty)
              _SduiActionBar(actions: actions, onAction: onAction),
          ],
        ),
      ),
    );
  }
}

class _SduiComponentNode extends StatelessWidget {
  const _SduiComponentNode({
    required this.component,
    required this.engine,
    required this.onAction,
  });

  final SduiComponent component;
  final SduiEngine engine;
  final ValueChanged<SduiAction>? onAction;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final items = _readItems(component.data['items']);
    final emphasis = component.data['emphasis'] as String?;
    final color = _surfaceColor(theme.colorScheme, emphasis);

    if (component.type == 'text') {
      return Text(
        component.body ?? component.title ?? '',
        style: theme.textTheme.bodyMedium,
      );
    }

    if (component.type == 'badge') {
      return Align(
        alignment: Alignment.centerLeft,
        child: Chip(
          label: Text(component.title ?? component.body ?? component.id),
        ),
      );
    }

    if (component.type == 'button') {
      final primaryAction =
          component.actions.where((action) => action.visible).isNotEmpty
          ? component.actions.where((action) => action.visible).first
          : null;
      return Align(
        alignment: Alignment.centerLeft,
        child: ElevatedButton(
          onPressed: primaryAction == null || onAction == null
              ? null
              : () => onAction!(primaryAction),
          child: Text(component.title ?? primaryAction?.label ?? 'Action'),
        ),
      );
    }

    if (!engine.supportsComponent(component)) {
      return _SduiContainer(
        color: theme.colorScheme.surfaceContainerHighest.withValues(
          alpha: 0.45,
        ),
        child: Text(
          'Unsupported component: ${component.type}',
          style: theme.textTheme.bodyMedium?.copyWith(
            color: theme.colorScheme.onSurfaceVariant,
          ),
        ),
      );
    }

    final content = _SduiContainer(
      color: color,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          if (component.title != null)
            Text(
              component.title!,
              style: theme.textTheme.titleMedium?.copyWith(
                fontWeight: FontWeight.w600,
              ),
            ),
          if (component.body != null) ...[
            if (component.title != null) const SizedBox(height: 6),
            Text(
              component.body!,
              style: theme.textTheme.bodyMedium?.copyWith(height: 1.35),
            ),
          ],
          if (items.isNotEmpty) ...[
            const SizedBox(height: 12),
            for (final item in items) ...[
              _SduiBulletRow(label: item),
              const SizedBox(height: 8),
            ],
          ],
          if (component.children
              .where((child) => child.visible)
              .isNotEmpty) ...[
            const SizedBox(height: 12),
            for (final child in component.children.where(
              (child) => child.visible,
            )) ...[
              _SduiComponentNode(
                component: child,
                engine: engine,
                onAction: onAction,
              ),
              const SizedBox(height: 12),
            ],
          ],
          if (component.actions.where((action) => action.visible).isNotEmpty)
            _SduiActionBar(
              actions: component.actions
                  .where((action) => action.visible)
                  .toList(growable: false),
              onAction: onAction,
            ),
        ],
      ),
    );

    if (component.type == 'card') {
      return Card(child: content);
    }

    if (component.type == 'list' || component.type == 'bullet_list') {
      return ListView(
        shrinkWrap: true,
        physics: const NeverScrollableScrollPhysics(),
        children: [content],
      );
    }

    return content;
  }
}

class _SduiContainer extends StatelessWidget {
  const _SduiContainer({required this.color, required this.child});

  final Color color;
  final Widget child;

  @override
  Widget build(BuildContext context) {
    return DecoratedBox(
      decoration: BoxDecoration(
        color: color,
        borderRadius: BorderRadius.circular(16),
      ),
      child: Padding(padding: const EdgeInsets.all(16), child: child),
    );
  }
}

class _SduiActionBar extends StatelessWidget {
  const _SduiActionBar({required this.actions, required this.onAction});

  final List<SduiAction> actions;
  final ValueChanged<SduiAction>? onAction;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 8,
      runSpacing: 8,
      children: [
        for (final action in actions)
          FilledButton.tonal(
            onPressed: onAction == null ? null : () => onAction!(action),
            child: Text(
              action.confirm ? '${action.label} (confirm)' : action.label,
            ),
          ),
      ],
    );
  }
}

class _SduiBulletRow extends StatelessWidget {
  const _SduiBulletRow({required this.label});

  final String label;

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.only(top: 6),
          child: Icon(Icons.circle, size: 8, color: theme.colorScheme.primary),
        ),
        const SizedBox(width: 10),
        Expanded(
          child: Text(
            label,
            style: theme.textTheme.bodyMedium?.copyWith(height: 1.35),
          ),
        ),
      ],
    );
  }
}

List<String> _readItems(Object? value) {
  if (value is! List) {
    return const <String>[];
  }

  return value.map((entry) => entry.toString()).toList(growable: false);
}

Color _surfaceColor(ColorScheme colorScheme, String? emphasis) {
  switch (emphasis) {
    case 'positive':
      return colorScheme.primaryContainer.withValues(alpha: 0.45);
    case 'caution':
      return colorScheme.tertiaryContainer.withValues(alpha: 0.5);
    default:
      return colorScheme.surfaceContainerHighest.withValues(alpha: 0.45);
  }
}
