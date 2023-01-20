import React from "react";
import { Snackbar } from "@mui/material";
import MuiAlert from '@mui/material/Alert';

const Alert = React.forwardRef(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });

export function MessageSucces({text}) {
    return (
    <Snackbar autoHideDuration={600}>
        <Alert severity="success" sx={{ width: '100%' }}>
            {text}
        </Alert>
    </Snackbar>
)}

export function MessageError({text}) {
    return (
    <Snackbar autoHideDuration={600}>
        <Alert severity="error" sx={{ width: '100%' }}>
            {text}
        </Alert>
    </Snackbar>
)}