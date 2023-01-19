import React, { Component } from 'react';
import Typography from '@mui/material/Typography';
import './index.css';
import { Grid, Container, Box} from '@mui/material';
import {Modal, Button} from '@mui/material';
import {CardContent, Card, CardMedia, CardActions} from '@mui/material';


class App extends Component {
  render() {
    return <Container>
      <Grid container spacing={2}>
        <Grid item xs={6}><Word/></Grid>
        <Grid item xs={3}>
          <ModalBtn button="Create new card"><NewCard/></ModalBtn>
        </Grid>
      </Grid>
    </Container>
  }
}

function ModalBtn({children={}, button={}}) {
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  return <section>    
    <Button  variant="outlined" onClick={handleOpen}>{button}</Button>
    <Modal open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
    >
      {children}
    </Modal>
  </section>
}

const NewCard = React.forwardRef((props, ref) => {
  return <Box className="modal">
      <Typography id="modal-modal-title" variant="h6" component="h2">
        Create a new card
      </Typography>
      <Typography id="modal-modal-description">
        Duis mollis, est non commodo luctus, nisi erat porttitor ligula.
      </Typography>
    </Box>
})

function Word({data = {value:"", description:"", lang:"", image: {hash: "test"}}}) {
  return (
    <Card sx={{ maxWidth: 340 }}>
      <CardMedia sx={{ height: 100 }}
        image={`/v1/image/{hash}`}
        title=""
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {value}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {description}
        </Typography>
      </CardContent>
      <CardActions>
        <Button size="small">Share</Button>
        <Button size="small">Learn More</Button>
      </CardActions>
    </Card>
  );
}

export default App;
