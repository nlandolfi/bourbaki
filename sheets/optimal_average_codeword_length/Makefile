# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make optimal_average_codeword_length.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

optimal_average_codeword_length.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimal_average_codeword_length.tex

optimal_average_codeword_length.pdf: ../../*.tex ../../trademark.pdf *.tex optimal_average_codeword_length.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode optimal_average_codeword_length.tex
	../../latexrun optimal_average_codeword_length.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimal_average_codeword_length.pdf
	open optimal_average_codeword_length.pdf

o:
	make open

terms: optimal_average_codeword_length.tex
	bbk_sheet -mode ts optimal_average_codeword_length.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f optimal_average_codeword_length.*
	make reset
	bbk_sheet -mode mk > Makefile