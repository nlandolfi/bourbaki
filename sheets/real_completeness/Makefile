# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_completeness.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_completeness.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_completeness.tex

real_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex real_completeness.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_completeness.tex
	../../latexrun real_completeness.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_completeness.pdf
	open real_completeness.pdf

o:
	make open

terms: real_completeness.tex
	bbk sheet -mode ts real_completeness.tex

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
	rm -f real_completeness.*
	make reset
	bbk sheet -mode mk > Makefile