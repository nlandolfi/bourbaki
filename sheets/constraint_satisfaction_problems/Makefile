# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make constraint_satisfaction_problems.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

constraint_satisfaction_problems.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > constraint_satisfaction_problems.tex

constraint_satisfaction_problems.pdf: ../../*.tex ../../trademark.pdf *.tex constraint_satisfaction_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode constraint_satisfaction_problems.tex
	../../latexrun constraint_satisfaction_problems.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: constraint_satisfaction_problems.pdf
	open constraint_satisfaction_problems.pdf

o:
	make open

terms: constraint_satisfaction_problems.tex
	bbk sheet -mode ts constraint_satisfaction_problems.tex

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
	rm -f constraint_satisfaction_problems.*
	make reset
	bbk sheet -mode mk > Makefile