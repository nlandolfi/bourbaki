# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make generalized_assignment_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

generalized_assignment_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > generalized_assignment_problems.tex

generalized_assignment_problems.pdf: ../../*.tex ../../trademark.pdf *.tex generalized_assignment_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode generalized_assignment_problems.tex
	../../latexrun generalized_assignment_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generalized_assignment_problems.pdf
	open generalized_assignment_problems.pdf

o:
	make open

terms: generalized_assignment_problems.tex
	bbk sheet -mode ts generalized_assignment_problems.tex

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
	rm -f generalized_assignment_problems.*
	make reset
	bbk sheet -mode mk > Makefile