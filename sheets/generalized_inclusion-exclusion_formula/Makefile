# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make generalized_inclusion-exclusion_formula.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

generalized_inclusion-exclusion_formula.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > generalized_inclusion-exclusion_formula.tex

generalized_inclusion-exclusion_formula.pdf: ../../*.tex ../../trademark.pdf *.tex generalized_inclusion-exclusion_formula.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode generalized_inclusion-exclusion_formula.tex
	../../latexrun generalized_inclusion-exclusion_formula.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generalized_inclusion-exclusion_formula.pdf
	open generalized_inclusion-exclusion_formula.pdf

o:
	make open

terms: generalized_inclusion-exclusion_formula.tex
	bbk_sheet -mode ts generalized_inclusion-exclusion_formula.tex

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
	rm -f generalized_inclusion-exclusion_formula.*
	make reset
	bbk_sheet -mode mk > Makefile