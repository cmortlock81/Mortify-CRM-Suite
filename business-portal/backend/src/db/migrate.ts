import fs from 'fs'; import path from 'path'; import {fileURLToPath} from 'url'; import {pool} from './pool.js';
const dir=path.resolve(path.dirname(fileURLToPath(import.meta.url)),'../../migrations');
await pool.query('create table if not exists schema_migrations(name text primary key, run_at timestamptz not null default now())');
for(const f of fs.readdirSync(dir).filter(f=>f.endsWith('.sql')).sort()){const done=await pool.query('select 1 from schema_migrations where name=$1',[f]); if(done.rowCount) continue; console.log('migrating',f); await pool.query('begin'); try{await pool.query(fs.readFileSync(path.join(dir,f),'utf8')); await pool.query('insert into schema_migrations(name) values($1)',[f]); await pool.query('commit')}catch(e){await pool.query('rollback'); throw e}}
await pool.end();
