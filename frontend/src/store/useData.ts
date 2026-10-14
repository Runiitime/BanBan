import React from "react";

import type { IBoard, ICard, IColumn } from "@data/board";
import type { IMap } from "@data/structure";

export const useData = () => {
  const boardInfo: IBoard = React.useMemo(
    () => ({
      id: "board_1",
      title: "Board 1",
      description: "A board",
      columnsOrder: ["clm", "clm2"],
    }),
    [],
  );

  const boardColumns: IMap<IColumn> = React.useMemo(
    () => ({
      clm: {
        id: "clm",
        boardId: "board_1",
        title: "First column",
        description: "---",
        position: 1,
        color: "",
        created_at: "---",
        updated_at: "---",
        itemsOrder: ["clm_crd1", "clm_crd2"],
      },
      clm2: {
        id: "clm2",
        boardId: "board_1",
        title: "Second column",
        description: "---",
        position: 1,
        color: "",
        created_at: "---",
        updated_at: "---",
        itemsOrder: ["clm2_crd3"],
      },
    }),
    [],
  );

  const columnsRows: IMap<ICard> = React.useMemo(
    () => ({
      clm_crd1: {
        id: "clm_crd1",
        position: 1,
        columnId: "clm",
        title: "First card",
        description: "---",
        created_at: "---",
        updated_at: "---",
        completed_at: "---",
        status: "in process",
      },
      clm_crd2: {
        id: "clm_crd2",
        position: 2,
        columnId: "clm",
        title: "Second card",
        description: "---",
        created_at: "---",
        updated_at: "---",
        completed_at: "---",
        status: "in process",
      },
      clm2_crd3: {
        id: "clm2_crd3",
        position: 3,
        columnId: "clm2",
        title: "Third card",
        description: "---",
        created_at: "---",
        updated_at: "---",
        completed_at: "---",
        status: "in process",
      },
    }),
    [],
  );

  const rowsOrder = React.useMemo(() => {
    const order: IMap<string[]> = {};

    boardInfo.columnsOrder.forEach((clm: string) => {
      const column = boardColumns[clm];

      order[clm] = column.itemsOrder;
    });

    return order;
  }, [boardInfo, boardColumns]);

  return { boardInfo, boardColumns, columnsRows, rowsOrder };
};
