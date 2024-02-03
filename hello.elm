module Example exposing (main)

import Html exposing (..)
import List
import Html.Attributes exposing (..)


hello name =
    div [] [ text <| "Hello, " ++ name ++ ", here are your TODOs" ]

todos items =
    ul [] (List.map todo items)

todo item =
    li [] [(case item.done of
        True ->
            text <| item.title ++ " - done"
        False ->
            text <| item.title ++ " - nope"
    )]

greeting data =
    div [ class "" ]
        [ (hello data.hello)
        , (todos data.todos)
        ]


main =
    greeting
    { hello = "Kristoffer"
    , todos =
        [ { done = True, title = "Shower" }
        , { done = True, title = "Drink (sparkling) water" }
        , { done = False, title = "Switch to Yerba Mate" }
        , { done = False, title = "Finish the fosdem lightning talk in time" }
        ]
    }
