# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make inequality_constrained_space_optimization_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

inequality_constrained_space_optimization_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > inequality_constrained_space_optimization_problems.tex

inequality_constrained_space_optimization_problems.pdf: ../../*.tex ../../trademark.pdf *.tex inequality_constrained_space_optimization_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode inequality_constrained_space_optimization_problems.tex
	../../latexrun inequality_constrained_space_optimization_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inequality_constrained_space_optimization_problems.pdf
	open inequality_constrained_space_optimization_problems.pdf

o:
	make open

terms: inequality_constrained_space_optimization_problems.tex
	bbk sheet -mode ts inequality_constrained_space_optimization_problems.tex

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
	rm -f inequality_constrained_space_optimization_problems.*
	make reset
	bbk sheet -mode mk > Makefile