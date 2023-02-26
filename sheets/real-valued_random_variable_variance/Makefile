# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real-valued_random_variable_variance.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real-valued_random_variable_variance.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real-valued_random_variable_variance.tex

real-valued_random_variable_variance.pdf: ../../*.tex ../../trademark.pdf *.tex real-valued_random_variable_variance.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real-valued_random_variable_variance.tex
	../../latexrun real-valued_random_variable_variance.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real-valued_random_variable_variance.pdf
	open real-valued_random_variable_variance.pdf

o:
	make open

terms: real-valued_random_variable_variance.tex
	bbk_sheet -mode ts real-valued_random_variable_variance.tex

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
	rm -f real-valued_random_variable_variance.*
	make reset
	bbk_sheet -mode mk > Makefile