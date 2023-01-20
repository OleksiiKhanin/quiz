import React from 'react';
import { Box, Typography, Button } from "@mui/material";
import { Photo } from "./photo"
import { WordInput } from "./word"
import { CreateNewCardPair } from './urls';
import {Languages} from './select';
import { objToImg, imgToObj } from "./photo";
import './index.css';

export const NewCard = React.forwardRef(({onError, onSuccess}, ref) =>{
    const [cardpair, setCardpair] = React.useState({
      cards: [
          {value: "", description: "", lang: Languages[0],},
          {value: "", description: "", lang: Languages[1],},
      ],
      image_type: "",
      image_data: "",
      image_title: "",
  })
  const loadCardPair = async ()=>{
    CreateNewCardPair({cardpair: cardpair, onError: onError, onSuccess: onSuccess})
  }
    return <Box className="modal">
        <Photo className="center" 
          src={objToImg({type: cardpair.image_type, data: cardpair.image_data})}
          title={cardpair.image_title}
          onLoad={(img,title)=>{
            const {data, type} = imgToObj(img) 
            setCardpair({...cardpair, image_data:data, image_type:type, image_title:title})}
          }
        />
        <Typography margin={"10px"} variant="h6" component="h2">Insert the word</Typography>
        <WordInput 
          word={cardpair.cards[0]} 
          onInput={word=>setCardpair({...cardpair, cards: [{...word},{...cardpair.cards[1]}]})}
        />
        <Typography margin={"10px"} variant="h6" component="h2">Insert the card backside</Typography>
        <WordInput 
          word={cardpair.cards[1]} 
          onInput={word=>setCardpair({...cardpair, cards: [{...cardpair.cards[0]},{...word}]})}
        />
        <Button variant="contained" onClick={loadCardPair}>Upload</Button>
      </Box>
  })