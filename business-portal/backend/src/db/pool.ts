import pg from 'pg'; import { env } from '../config/env.js';
export const pool=new pg.Pool({connectionString:env.databaseUrl});
export async function q<T=any>(text:string,params:any[]=[]){const r=await pool.query<T>(text,params); return r;}
