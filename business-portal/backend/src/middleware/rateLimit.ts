import rateLimit from 'express-rate-limit'; export const authRateLimit=rateLimit({windowMs:15*60*1000,limit:30}); export const agentRateLimit=rateLimit({windowMs:60*1000,limit:120});
