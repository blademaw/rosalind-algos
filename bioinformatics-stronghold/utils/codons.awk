#!/usr/bin/awk -f

{
    # Process each line, extracting codon-amino acid pairs
    for (i = 1; i <= NF; i += 2) {
        if (i + 1 <= NF) {
            codon = $i
            aa = $(i + 1)
            
            # Handle special case for Stop codons
            if (aa == "Stop") {
                printf "{\"" codon "\", '*'},"
            } else {
                printf "{\"" codon "\", '" aa "'},"
            }
            
            # Add space between pairs for readability
            if (i + 2 <= NF) printf " "
        }
    }
    printf "\n"
}

END {
    # Remove the trailing comma from the last line
    printf "\n"
}
