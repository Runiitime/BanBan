import React from "react";

import type { IColumn } from "@data/board";
import type { IMap } from "@data/structure";

export const useItemsOrder = (items: IMap<IColumn>) => {
  const columnOrders: string[] = React.useMemo(() => {
    return Object.keys(items);
  }, [items]);

  const columnItemsOrders = React.useMemo(() => {
    const itemsOrder: IMap<string[]> = {};

    columnOrders.forEach((clm: string) => {
      const column: IColumn = items[`${clm}`];
      itemsOrder[clm] = Object.keys(column.cards);
    });

    return itemsOrder;
  }, [items, columnOrders]);

  return { columnOrders, columnItemsOrders };
};
