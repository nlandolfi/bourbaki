# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make maximum_conditional_estimates.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

maximum_conditional_estimates.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > maximum_conditional_estimates.tex

maximum_conditional_estimates.pdf: ../../*.tex ../../trademark.pdf *.tex maximum_conditional_estimates.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode maximum_conditional_estimates.tex
	../../latexrun maximum_conditional_estimates.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: maximum_conditional_estimates.pdf
	open maximum_conditional_estimates.pdf

o:
	make open

terms: maximum_conditional_estimates.tex
	bbk sheet -mode ts maximum_conditional_estimates.tex

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
	rm -f maximum_conditional_estimates.*
	make reset
	bbk sheet -mode mk > Makefile