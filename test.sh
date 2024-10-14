          #!/bin/bash
          INPUT_DESTINATION="committer"
          EMAIL="jarcher@patriotsoftware.com"

          if [[ "$INPUT_DESTINATION" == "#"* ]]; then
            echo "1"
         
          elif [[ "$INPUT_DESTINATION" == "committer" && "$EMAIL" = *'@patriotsoftware.com' ]]; then
            echo "2"
            
          else 
            echo "3"
           
          fi