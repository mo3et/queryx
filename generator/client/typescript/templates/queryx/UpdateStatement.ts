import { Clause } from "./Clause";

export const newUpdate = (table: string) => {
  return new UpdateStatement(table);
};

export class UpdateStatement {
  private _table: string;
  private _columns?: string[];
  private _values?: any[];
  private _where?: Clause;
  private _returning?: string[];

  constructor(table: string) {
    this._table = table;
  }

  columns(...columns: string[]) {
    this._columns = columns;
    return this;
  }

  values(...values: any[]) {
    this._values = values;
    return this;
  }

  where(expr: Clause) {
    this._where = expr;
    return this;
  }

  returning(...returning: string[]) {
    this._returning = returning;
    return this;
  }

  toSQL(): [string, any[]] {
    let sql = `UPDATE ${this._table} SET`;
    let args: any[] = [];

    if (this._values !== undefined) {
      args = args.concat(this._values);
    }

    let sets = [];
    if (this._columns !== undefined) {
      for (let col of this._columns) {
        sets.push(`${col} = ?`);
      }
    }

    sql = `${sql} ${sets.join(", ")}`;

    if (this._where !== undefined) {
      sql = `${sql} WHERE ${this._where.fragment}`;
      args = args.concat(this._where.args);
    }

    if (this._returning !== undefined) {
      sql = `${sql} RETURNING ${this._returning.join(", ")}`;
    }

    return [sql, args];
  }
}
