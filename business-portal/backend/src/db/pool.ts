import pg, {type QueryResultRow} from 'pg'; import { env } from '../config/env.js';
export const pool=new pg.Pool({connectionString:env.databaseUrl});
export async function q<T extends QueryResultRow=QueryResultRow>(text:string,params:any[]=[]){const r=await pool.query<T>(text,params); return r;}
