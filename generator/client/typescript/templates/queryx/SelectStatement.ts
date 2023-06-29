import { Clause } from "./Clause";
import { DeleteStatemnet, newDelete } from "./DeleteStatement";
import { newUpdate } from "./UpdateStatement";

export const newSelect = () => {
  return new SelectStatement();
};

export class SelectStatement {
  private _selection?: string[];
  private _from?: string;
  private _where?: Clause;
  private _limit?: number;
  private _offset?: number;
  private _order: string[] = [];
  private _group?: string;
  private _having?: string;
  private _joins: string[] = [];

  select(...selection: string[]) {
    this._selection = selection;
    return this;
  }

  from(from: string) {
    this._from = from;
    return this;
  }

  where(expr: Clause) {
    this._where = expr;
    return this;
  }

  limit(limit: number) {
    this._limit = limit;
    return this;
  }

  offset(offset: number) {
    this._offset = offset;
    return this;
  }

  order(...order: string[]) {
    this._order = [...this._order, ...order];
    return this;
  }

  groupBy(group: string) {
    this._group = group;
    return this;
  }

  having(having: string) {
    this._having = having;
    return this;
  }

  // convert select statement to update statement
  update() {
    let s;
    if (this._limit != undefined) {
      // TODO: convert into subquery if limit
      s = newUpdate(this._from);
    } else {
      s = newUpdate(this._from);
    }

    if (this._where) {
      s.where(this._where);
    }

    return s;
  }

  // convert select statement to delete statement
  delete() {
    let s: DeleteStatemnet = null;

    if (this._from !== undefined) {
      s = newDelete(this._from);
    }

    if (this._where !== undefined) {
      s.where(this._where);
    }

    return s;
  }

  toSQL(): [string, any[]] {
    let query = "SELECT";
    let args = [];

    if (this._selection != undefined && this._selection.length > 0) {
      query += ` ${this._selection.join(", ")}`;
    }

    if (this._from !== undefined) {
      query += ` FROM ${this._from}`;
    }

    if (this._joins.length > 0) {
      for (let i = 0; i < this._joins.length; i++) {
        query += ` ${this._joins[i]}`;
      }
    }

    if (this._where !== undefined) {
      query += ` WHERE ${this._where.fragment}`;
      args.push(...this._where.args);
    }

    if (this._order.length > 0) {
      query += ` ORDER BY ${this._order.join(", ")}`;
    }

    if (this._limit !== undefined) {
      query += " LIMIT ?";
      args.push(this._limit);
    }

    if (this._offset !== undefined) {
      query += " OFFSET ?";
      args.push(this._offset);
    }

    if (this._group !== undefined) {
      query += ` GROUP BY ${this._group}`;
    }

    if (this._having !== undefined) {
      query += ` HAVING ${this._having}`;
    }

    return [query, args];
  }
}
