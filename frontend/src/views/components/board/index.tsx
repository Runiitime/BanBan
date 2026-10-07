import React from "react";
import { Box, Flex } from "@radix-ui/themes";
import { map } from "lodash";

import type { IBoard, IColumn } from "@data/board";

import { Column } from "../column";
import * as Styled from "./style";

interface Props {
  board: IBoard;
}

export const Board: React.FC<Props> = ({ board }: Props) => {
  const { columns } = board;

  const columnsItems = React.useMemo(() => {
    return map(columns, (item: IColumn) => <Column data={item} />);
  }, [columns]);

  return (
    <Box pt="4" px="7" style={Styled.Board}>
      <Flex direction="row" gap="4">
        {columnsItems}
      </Flex>
    </Box>
  );
};
