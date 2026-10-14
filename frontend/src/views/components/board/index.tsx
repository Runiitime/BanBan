import React from "react";
import { Card } from "@components/card";
import { move } from "@dnd-kit/helpers";
import {
  DragDropProvider,
  type DragEndEvent,
  type DragOverEvent,
} from "@dnd-kit/react";
import { Box, Flex } from "@radix-ui/themes";

import { useData } from "../../../store";
import { Column } from "../column";
import * as Styled from "./style";

export const Board: React.FC = () => {
  const { boardInfo, boardColumns, rowsOrder } = useData();

  const [columnOrder, setColumnOrder] = React.useState(
    () => boardInfo.columnsOrder,
  );

  const [itemsOrder, setItemsOrder] = React.useState(() => rowsOrder);

  // Move item
  const handleDragOver = React.useCallback((event: DragOverEvent) => {
    const { source, target } = event.operation;

    if (source?.type === "column") {
      return;
    }
    //
    setItemsOrder((items) => move(items, event));
  }, []);

  // Move column
  const handleDragEnd = React.useCallback((event: DragEndEvent) => {
    const { source, target } = event.operation;

    if (event.canceled || source?.type !== "column") {
      return;
    }

    setColumnOrder((columns) => move(columns, event));
  }, []);

  const columnsItems = React.useMemo(() => {
    return columnOrder.map((column: string, index: number) => {
      const clm = boardColumns[column];
      return (
        <Column key={clm?.id} index={index} {...clm}>
          {itemsOrder[clm?.id].map((id: string, index: number) => (
            <Card key={id} index={index} column={column} id={id} />
          ))}
        </Column>
      );
    });
  }, [boardColumns, columnOrder, itemsOrder]);

  return (
    <DragDropProvider onDragEnd={handleDragEnd} onDragOver={handleDragOver}>
      <Box pt="4" px="7" style={Styled.Board} id={"board"} className="Root">
        <Flex direction="row" gap="4">
          {columnsItems}
        </Flex>
      </Box>
    </DragDropProvider>
  );
};
