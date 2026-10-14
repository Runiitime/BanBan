import React, { useMemo } from "react";
import type { UniqueIdentifier } from "@dnd-kit/core";
import { useSortable } from "@dnd-kit/react/sortable";
import {
  Badge,
  Box,
  Card as CardUI,
  Flex,
  Heading,
  Text,
} from "@radix-ui/themes";

import { useData } from "../../../store";
import * as Styled from "./style.ts";

interface Props {
  id: UniqueIdentifier;
  index: number;
  column: string;
}

export const Card: React.FC<Props> = ({ id, column, index }: Props) => {
  const { columnsRows } = useData();

  const data = useMemo(() => {
    return columnsRows[id];
  }, [columnsRows, id]);

  const { ref, isDragging } = useSortable({
    id,
    index,
    type: "item",
    accept: "item",
    group: column,
  });

  return (
    <CardUI
      style={Styled.Card}
      className="Item"
      ref={ref}
      data-dragging={isDragging}
    >
      <Box>
        <Heading size="4">{data?.title}</Heading>
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
