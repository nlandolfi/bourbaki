# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make invariant_real_subspaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

invariant_real_subspaces.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > invariant_real_subspaces.tex

invariant_real_subspaces.pdf: ../../*.tex ../../trademark.pdf *.tex invariant_real_subspaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode invariant_real_subspaces.tex
	../../latexrun invariant_real_subspaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: invariant_real_subspaces.pdf
	open invariant_real_subspaces.pdf

o:
	make open

terms: invariant_real_subspaces.tex
	bbk_sheet -mode ts invariant_real_subspaces.tex

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
	rm -f invariant_real_subspaces.*
	make reset
	bbk_sheet -mode mk > Makefile