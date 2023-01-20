import axios from 'axios'

const GetCardByID = "/v1/card/id" // {id}
const GetCardByLang = "/v1/card/lang" // {lang}
const GetImageData = "/v1/image/data" // {hash}
const GetImageObject = "/v1/image/object" // {hash}

const CreateCardPair = "/v1/card" // post request
const CreateImageURL = "/v1/image" // post request


export async function CreateNewCardPair({cardpair = {
    cards: [
        {value: "", descriptio: "", lang: "english",},
        {value: "", descriptio: "", lang: "russian",},
    ],
    image_data: "",
    image_tittle: "",
}, onError=console.err, onSuccess=console.log}) {
    try {
        const answer = await axios.post(CreateCardPair, cardpair)
        onSuccess && onSuccess(answer.data)
    } catch (err) {
        onError && onError(err)
    }
}

export async function CreateImage({image = {
    tittle: "",
    data: "",
},onError=console.err, onSuccess=console.log}) {
    try {
        const answer = await axios.post(CreateImageURL, image)
        onSuccess && onSuccess(answer.data)
    } catch (err) {
        onError && onError(err)
    }
}

export async function GetImagesData({hash="",onError=console.err, onSuccess=console.log}) {
    try {
        const answer = await axios.get(`${GetImageData}/${hash}`)
        onSuccess && onSuccess(answer.data)
    } catch (err) {
        onError && onError(err)
    }
}

export async function GetImagesObject({hash="", onError=console.error, onSuccess=console.log}) {
    try {
        const answer = await axios.get(`${GetImageObject}/${hash}`)
        onSuccess && onSuccess(answer.data)
    } catch (err) {
        onError && onError(err)
    }
}

export async function GetRandomCard({lang="english", onError=console.err, onSuccess=console.log}) {
    try {
        const answer = await axios.get(`${GetCardByLang}/${lang}`)
        onSuccess && onSuccess(answer.data)
    } catch (err) {
        onError && onError(err)
    }
}

export async function GetCard({id=0, onError=console.err, onSuccess=console.log}) {
    try {
        const answer = await axios.get(`${GetCardByID}/${id}`)
        onSuccess && onSuccess(answer.data)
    } catch (err) {
        onError && onError(err)
    }
}