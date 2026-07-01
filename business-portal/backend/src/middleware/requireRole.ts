import {Request,Response,NextFunction} from 'express'; import {can,Module} from '../config/roles.js'; import {ApiError} from './errorHandler.js';
export const requirePermission=(module:Module,write=false)=>(req:Request,res:Response,next:NextFunction)=>{if(!req.user||!can(req.user.role,module,write)) return next(new ApiError(403,'FORBIDDEN','Insufficient role permission')); next()};
