# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_integral_limit_theorems.pdf

real_integral_limit_theorems.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_integral_limit_theorems.tex

real_integral_limit_theorems.pdf: ../../*.tex ../../trademark.pdf *.tex real_integral_limit_theorems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_integral_limit_theorems.tex
	../../latexrun real_integral_limit_theorems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_integral_limit_theorems.pdf
	open real_integral_limit_theorems.pdf

o:
	make open

terms: real_integral_limit_theorems.tex
	bbk_sheet -mode ts real_integral_limit_theorems.tex

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
	rm -f real_integral_limit_theorems.*
	make reset
	bbk_sheet -mode mk > Makefile