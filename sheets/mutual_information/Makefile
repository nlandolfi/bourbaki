# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make mutual_information.pdf

mutual_information.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > mutual_information.tex

mutual_information.pdf: ../../*.tex ../../trademark.pdf *.tex mutual_information.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode mutual_information.tex
	../../latexrun mutual_information.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: mutual_information.pdf
	open mutual_information.pdf

o:
	make open

terms: mutual_information.tex
	bbk_sheet -mode ts mutual_information.tex

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
	rm -f mutual_information.*
	make reset
	bbk_sheet -mode mk > Makefile