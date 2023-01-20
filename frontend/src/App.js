import React, { Component } from 'react';
import Typography from '@mui/material/Typography';
import { Grid, Container} from '@mui/material';
import {Modal, Button} from '@mui/material';
import {CardContent, Card, CardMedia, CardActionArea} from '@mui/material';
import { NewCard } from './NewCard';
import './index.css';
import { GetImagesObject, GetRandomCard } from './urls';
import { objToImg } from './photo';
import { Languages } from './select';

class App extends Component {
  render() {
    return <Container>
      <Grid container spacing={2}>
        <Grid item xs={2}>
          <ModalBtn button="Create card"><NewCard/></ModalBtn>
        </Grid>
      </Grid>
      <Cards/>
    </Container>
  }
}

function Cards() {
  const [cards, setCards] = React.useState([{}, {}])
  React.useEffect(()=>{
    GetRandomCard({lang: Languages[0], onSuccess:data=>setCards([...data])})
  }, [])
  return (
    <Grid container spacing={2}>
      <Grid item xs={3}>
        <Word origin={cards[0]} onFlip={()=>setCards([cards[1], cards[0]])}/>
      </Grid>
  </Grid>
  )
}

function ModalBtn({children={}, button={}}) {
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  return <section>    
    <Button  variant="outlined" onClick={handleOpen}>{button}</Button>
    <Modal open={open} onClose={handleClose}>
      {children}
    </Modal>
  </section>
}

function Word({origin = {value:"", description:"", lang:"", image_hash: ""}, onFlip=e=>console.log("Flipped")}) {
  const [image, setImage] = React.useState({data: "", title: ""})
  React.useEffect(()=>{
    if (origin && origin.image_hash && origin.image_hash.length > 0) {
      GetImagesObject({
        hash: origin.image_hash, 
        onSuccess: img=>setImage({data: objToImg({type: img.type,data:img.data}), title: img.title}),
        onError: err=>console.log(err),
      })
    }
  }, [origin.image_hash])
  return (
    <Card sx={{maxWidth: "100%"}}>
      <CardActionArea onClick={(e)=>onFlip && onFlip(e)}>
      <CardMedia 
        sx={{maxHeight: 450}}
        height="100%"
        component='img'
        image={image.data}
        title={image.title}
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {origin.value}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          {origin.description}
        </Typography>
      </CardContent>
      </CardActionArea>
    </Card>
  );
}

export default App;
