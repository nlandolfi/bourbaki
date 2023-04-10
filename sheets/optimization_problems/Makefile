# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make optimization_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

optimization_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > optimization_problems.tex

optimization_problems.pdf: ../../*.tex ../../trademark.pdf *.tex optimization_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode optimization_problems.tex
	../../latexrun optimization_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimization_problems.pdf
	open optimization_problems.pdf

o:
	make open

terms: optimization_problems.tex
	bbk sheet -mode ts optimization_problems.tex

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
	rm -f optimization_problems.*
	make reset
	bbk sheet -mode mk > Makefile