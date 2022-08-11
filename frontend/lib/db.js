import mysql from "serverless-mysql";

export const db = mysql({
  config: {
    host: process.env.MYSQL_HOST,
    database: process.env.MYSQL_DATABASE,
    user: process.env.MYSQL_USERNAME,
    password: process.env.MYSQL_PASSWORD.trim(),
  },
});

export async function sql_query(query_string) {
  try {
    const results = await db.query(query_string);
    await db.end();
    return results;
  } catch (e) {
    throw Error(e.message);
  }
}
