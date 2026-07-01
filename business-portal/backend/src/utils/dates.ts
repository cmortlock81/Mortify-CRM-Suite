export const now=()=>new Date().toISOString(); export const daysFromNow=(d:number)=>new Date(Date.now()+d*864e5).toISOString();
