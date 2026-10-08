import React from "react";
import type { UniqueIdentifier } from "@dnd-kit/core";
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import {
  Badge,
  Box,
  Card as CardUI,
  Flex,
  Heading,
  Text,
} from "@radix-ui/themes";

import type { ICard } from "@data/board";

import * as Styled from "./style.ts";

interface Props {
  id: UniqueIdentifier;
  data: ICard;
}

export const Card: React.FC<Props> = ({ id, data }: Props) => {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({
    id: id,
  });

  const styled: React.CSSProperties = {
    transform: CSS.Transform.toString(transform),
    transition,
    boxSizing: "border-box",
    backgroundColor: isDragging ? "#FF977D" : "#70B8FF",
    ...Styled.Card,
  };

  return (
    <CardUI {...attributes} {...listeners} style={styled} ref={setNodeRef}>
      <Box>
        <Heading size="4">{data.title}</Heading>
      </Box>
      <Box>
        <Text>{data?.description}</Text>
      </Box>
      <Box>
        <Flex gap="2">
          <Badge color="ruby">Test</Badge>
          <Badge color="bronze">Tdse</Badge>
          <Badge color="amber">Vde</Badge>
        </Flex>
      </Box>
    </CardUI>
  );
};
