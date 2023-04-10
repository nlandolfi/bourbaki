# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_function_epigraphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_function_epigraphs.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > real_function_epigraphs.tex

real_function_epigraphs.pdf: ../../*.tex ../../trademark.pdf *.tex real_function_epigraphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_function_epigraphs.tex
	../../latexrun real_function_epigraphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_function_epigraphs.pdf
	open real_function_epigraphs.pdf

o:
	make open

terms: real_function_epigraphs.tex
	bbk_sheet -mode ts real_function_epigraphs.tex

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
	rm -f real_function_epigraphs.*
	make reset
	bbk_sheet -mode mk > Makefile