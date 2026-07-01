import {Navigate,Outlet} from 'react-router-dom'; export const ProtectedRoute=()=>localStorage.getItem('token')?<Outlet/>:<Navigate to="/login"/>;
