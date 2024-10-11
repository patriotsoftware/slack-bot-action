import sys

def newResultLine(jobresult):
    
    match(jobresult.split(':')[1]):
        case "success":
            return f"✅ {jobresult.split(':')[0]} Succeeded.\n"
        case "failure":
            return f"❌ {jobresult.split(':')[0]} Failed. \n"
        case _:
            return f"❕ {jobresult.split(':')[0]} Didn't Run. \n"
           
def format_results(inputString):
    jobresult_list = inputString.splitlines()
   # formatted_list = [newResultLine(result) for result in inputString.splitlines()]
    #return ''.join(formatted_list)
    #return jobresult_list
    return "Test successful"

format_results()

