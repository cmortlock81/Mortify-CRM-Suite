export const SearchBox=({onSearch}:{onSearch:(q:string)=>void})=><input placeholder="Search" onChange={e=>onSearch(e.target.value)} />;
