# cef2xml_v2
A cef to xml converter for the CAA team.


# Hello, World!

go install github.com\caa-dev-apps\cef2xml_v2

set GOOS=linux 
set GOOS=darwin
set GOARCH=amd64
set GOARCH=386 

go install 




        !
        FILE_NAME = "C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01.cef"
        !
        include =   "C1_CH_PEA_PITCH_3DXLARH_DPFlux.ceh" ! Contains all the static meta-data
        !
        ! Start of non-Static File Level - i.e. can't be from included file
        START_META = LOGICAL_FILE_ID
            ENTRY   = "C1_CP_PEA_PITCH_3DXLARH_DPFlux__20111031_V01" ! Filename less the .cef extension
        END_META   = LOGICAL_FILE_ID
        !






