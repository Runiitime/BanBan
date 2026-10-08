import React from "react";
import {
  closestCenter,
  DndContext,
  type DragCancelEvent,
  type DragEndEvent,
  DragOverlay,
  type DragStartEvent,
  KeyboardSensor,
  PointerSensor,
  type UniqueIdentifier,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { Box, Flex, Heading } from "@radix-ui/themes";

import type { ICard, IColumn } from "@data/board";

import { Card } from "../card";
import * as Styled from "./style";

interface Props {
  data: IColumn;
}

export const Column: React.FC<Props> = ({ data }: Props) => {
  const [items, setItems] = React.useState(data.cards);
  const [activeId, setActiveId] = React.useState<UniqueIdentifier | null>(null);
  const [activeItem, setActiveItem] = React.useState<ICard>();

  const { title } = data;

  const sortableCards = React.useMemo(() => {
    return items.map((item: ICard) => item.id);
  }, [items]);

  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: { delay: 100, tolerance: 5 },
    }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    }),
  );

  const handleDragStart = React.useCallback(
    (event: DragStartEvent) => {
      setActiveId(event.active.id);
      const found = items.find((item) => item.id === event.active.id);
      setActiveItem(found);
    },
    [items],
  );

  const handleDragEnd = React.useCallback((event: DragEndEvent) => {
    setActiveId(null);
    setActiveItem(undefined);
    const { active, over } = event;

    if (!over) {
      return;
    }

    if (active.id !== over.id) {
      setItems((items) => {
        const oldIdx = items.findIndex((item) => item.id === active.id);
        const newIdx = items.findIndex((item) => item.id === over.id);

        return arrayMove(items, oldIdx, newIdx);
      });
    }
  }, []);

  const handleDragCancel = React.useCallback((event: DragCancelEvent) => {
    void event;
    setActiveId(null);
    setActiveItem(undefined);
  }, []);

  return (
    <Box style={Styled.Column}>
      <Box py="3" mb="4" style={Styled.Heading}>
        <Heading>{title}</Heading>
      </Box>
      <Flex direction={"column"} gap={"3"}>
        <DndContext
          sensors={sensors}
          collisionDetection={closestCenter}
          onDragStart={handleDragStart}
          onDragEnd={handleDragEnd}
          onDragCancel={handleDragCancel}
        >
          <SortableContext
            items={sortableCards}
            strategy={verticalListSortingStrategy}
          >
            <Flex gap={"4"} direction={"column"}>
              {items.map((item: ICard) => (
                <Card key={item.id} id={item.id} data={item} />
                // <Item key={item.id} id={item.id} data={item} />
              ))}
            </Flex>
          </SortableContext>
          <DragOverlay
            adjustScale
            dropAnimation={{
              duration: 150,
              easing: `cubic-bezier(0.18, 0.67, 0.6, 1.22)`,
            }}
          >
            {activeId ? (
              <Card
                key={activeItem?.id}
                id={activeItem?.id as UniqueIdentifier}
                data={activeItem as ICard}
              />
            ) : null}
          </DragOverlay>
        </DndContext>
      </Flex>
    </Box>
  );
};
