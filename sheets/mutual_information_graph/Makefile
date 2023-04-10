# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make mutual_information_graph.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

mutual_information_graph.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > mutual_information_graph.tex

mutual_information_graph.pdf: ../../*.tex ../../trademark.pdf *.tex mutual_information_graph.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode mutual_information_graph.tex
	../../latexrun mutual_information_graph.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: mutual_information_graph.pdf
	open mutual_information_graph.pdf

o:
	make open

terms: mutual_information_graph.tex
	bbk_sheet -mode ts mutual_information_graph.tex

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
	rm -f mutual_information_graph.*
	make reset
	bbk_sheet -mode mk > Makefile