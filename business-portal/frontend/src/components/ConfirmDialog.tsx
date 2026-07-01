export const ConfirmDialog=({message,onConfirm}:{message:string,onConfirm:()=>void})=><button onClick={()=>confirm(message)&&onConfirm()}>Confirm</button>;
