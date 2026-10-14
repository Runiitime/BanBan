import React from "react";
import { CollisionPriority } from "@dnd-kit/abstract";
import { useSortable } from "@dnd-kit/react/sortable";
import { Box, Flex, Heading } from "@radix-ui/themes";

import type { IColumn } from "@data/board";

import * as Styled from "./style";

interface Props extends IColumn, React.PropsWithChildren {
  id: string;
  index: number;
}

export const Column: React.FC<Props> = (props: Props) => {
  const { id, index, title, color, description, children } = props;

  const { ref } = useSortable({
    id,
    index,
    type: "column",
    collisionPriority: CollisionPriority.Low,
    accept: ["item", "column"],
  });

  return (
    <Box style={Styled.Column} ref={ref}>
      <Box py="3" mb="4" style={Styled.Heading}>
        <Heading>{title}</Heading>
      </Box>
      <Flex direction={"column"} gap={"3"}>
        {children}
      </Flex>
    </Box>
  );
};
