# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_function_approximators.pdf

real_function_approximators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_function_approximators.tex

real_function_approximators.pdf: ../../*.tex ../../trademark.pdf *.tex real_function_approximators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_function_approximators.tex
	../../latexrun real_function_approximators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_function_approximators.pdf
	open real_function_approximators.pdf

o:
	make open

terms: real_function_approximators.tex
	bbk_sheet -mode ts real_function_approximators.tex

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
	rm -f real_function_approximators.*
	make reset
	bbk_sheet -mode mk > Makefile